import Logger from './Logger';
// @ts-ignore
import * as SpotifyWebApi from 'spotify-web-api-node';

export default class APIManager {
  public spotify: SpotifyWebApi;

  public constructor() {
    this.spotify = new SpotifyWebApi({
      clientId: process.env.SPOTIFY_ID,
      clientSecret: process.env.SPOTIFY_SECRET
    });
  }

  public async init() {
    try {
      const { body } = await this.spotify.clientCredentialsGrant();
      this.spotify.setAccessToken(body.access_token);
    } catch (e) {
      Logger.error(e);
    }
  }
}
