// @ts-ignore
import * as SpotifyWebApi from 'spotify-web-api-node';
import JunkoClient from '../client/JunkoClient';

export class APIManager {
  public spotify = new SpotifyWebApi({
    clientId: process.env.SPOTIFY_ID,
    clientSecret: process.env.SPOTIFY_SECRET
  });

  constructor(private readonly client: JunkoClient) {}

  public async init() {
    try {
      const { body } = await this.spotify.clientCredentialsGrant();
      this.spotify.setAccessToken(body.access_token);
    } catch (e) {
      this.client.logger.error(e);
    }
  }
}
