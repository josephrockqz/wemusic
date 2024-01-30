## Structure of Backend Application

Application structure is based on the following:
- [golang backend developer guide](https://medium.com/geekculture/how-to-structure-your-project-in-golang-the-backend-developers-guide-31be05c6fdd9)
- [golang gin API example](https://github.com/eddycjy/go-gin-example/blob/master/pkg/setting/setting.go)

## Prerequisites

Do at least one of the following in order for Spotify requests to work:
- Update `SPOTIFY_CLIENT_ID` & `SPOTIFY_CLIENT_SECRET` environment variables in the local-template.yaml config file
- Export Spotify application credentials in order to make requests with the following commands:
`export SPOTIFY_CLIENT_ID=...`
`export SPOTIFY_CLIENT_SECRET=...`
