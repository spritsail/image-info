const request = require('request-promise')
const errors = require('request-promise/errors');
const moment = require('moment')

const path = require('path')
const util = require('util')
let badgeSync = require('gh-badges')
badgeSync[util.promisify.custom] = (format) => {
  return new Promise((resolve, reject) => {
    badgeSync(format, (svg, err) => {
      if (err)
        reject(err)
      else
        resolve(svg)
    });
  });
};
const badge = util.promisify(badgeSync)

badgeSync.loadFont('./Verdana.ttf', err => {
  if (err) {
    console.log(err)
    process.exit(1)
  }
})

var app = require('express')()

app.use('/', (req, res, next) => {
  console.log(`[${new Date().toISOString()}]`, req.ip, req.method, req.originalUrl)
  next()
})

app.all('/info/:owner/:image', (req, res) => {
  let mbapi_url = 'https://api.microbadger.com/v1'
  let image = `${req.params.owner}/${req.params.image}`
  let url = `${mbapi_url}/images/${image}`

  request({uri: url, json: true})
    .then(json => {
      res.json({
        image: json['ImageName'],
        tag: json['LatestVersion'],
        tags: json['Versions'].map(ver => ({
          tags: ver['Tags'].map(t => t['tag']),
          last: ver['Created'],
          ago: moment(ver['Created']).fromNow()
        })),
        last: json['LastUpdated'],
        ago: moment(json['LastUpdated']).fromNow()
      })
    })
    .catch(errors.StatusCodeError, err => res.status(err.statusCode).send({error: err}))
    .catch(err => res.status(500).send({error: 'Internal Server Error'}))
})

app.all('/lastbuild/:owner/:image.svg', (req, res) => {
  let owner = req.params.owner
  let image = req.params.image.split(':')[0]
  let tag   = req.params.image.split(':')[1] || 'latest'

  let badgefmt = (arg) => badge({
    text: [
      req.query.text || 'last build',
      arg.msg || moment(arg.date).fromNow() || 'error'
    ],
    colorscheme: arg.color || req.query.color || 'blue',
    template: req.query.template || 'flat'
  })

  request.get({uri: `https://hub.docker.com/v2/repositories/${owner}/${image}/tags/`, json: true})
    .then(resp => {
      let tags = resp['results'].filter(t => t && (t.name == tag))
      if (tags.length < 1)
        return badgefmt({msg: 'no such tag', color: 'red'})
      else
        return badgefmt({date: tags[0]['last_updated']})
    })
    .catch(errors.StatusCodeError, err => {
      if (err && err.statusCode == 404)
        return badgefmt({msg: 'not found', color: 'red'})
      else
        throw err
    })
    .then(data => res.set('Content-Type', 'image/svg+xml').send(data))
    .catch(errors.RequestError, err => res.status(err.status).send(err))
    .catch(errors.TransformError, err => res.status(err.status).send(err))
    .catch(err => res.sendStatus(500))
    .finally(res.end.bind(res))
})
app.all('/lastbuild/:owner/:image', (req, res) => res.redirect(req.path + '.svg'))

var port = process.env.port || 8080
app.listen(port, _ => console.log(
  `[${new Date().toISOString()}]`,
  `${path.basename(process.argv[1])}:`,
  'Listening on', port)
)
