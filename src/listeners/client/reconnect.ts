import { Listener } from 'discord-akairo';
import { Logger } from '../../structs/Logger';

export default class ReconnectListener extends Listener {
  public constructor() {
    super('reconnect', {
      emitter: 'client',
      event: 'reconnect'
    });
  }

  public exec(): void {
    Logger.info('Reconnecting...');
  }
}
