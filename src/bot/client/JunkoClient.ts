import { Message } from 'discord.js';
import { AkairoClient, CommandHandler, InhibitorHandler, ListenerHandler } from 'discord-akairo';
import { Connection } from 'typeorm';
import { join } from 'path';
import Database from '../structs/Database';
import { Settings } from '../models/Settings';
import SettingsProvider from '../structs/SettingsProvider';
import Prometheus from '../structs/Prometheus';
import RPCServer from '../structs/RPCServer';
import APIManager from '../structs/APIManager';
import ReplyManager from '../structs/ReplyMenager';

interface JunkoConf {
  ownerID: string;
  token: string;
  color: string;
  defaultPrefix: string;
  defaultPreset: string;
}

declare module 'discord-akairo' {
  interface AkairoClient {
    config: JunkoConf;
    db: Connection;
    settings: SettingsProvider;
    prometheus: Prometheus;
    rpc: RPCServer;
    replyManager: ReplyManager;
    APIManager: APIManager;
    commandHandler: CommandHandler;
    inhibitorHandler: InhibitorHandler;
    listenerHandler: ListenerHandler;
  }
}

export default class JunkoClient extends AkairoClient {
  public constructor(config: JunkoConf) {
    super(
      {
        ownerID: config.ownerID
      },
      {
        messageCacheMaxSize: 50,
        messageCacheLifetime: 60 * 15,
        messageSweepInterval: 60 * 45,
        disableEveryone: true,
        disabledEvents: ['TYPING_START']
      }
    );

    this.config = config;

    this.prometheus = new Prometheus();

    this.rpc = new RPCServer(this);

    this.replyManager = new ReplyManager(this);

    this.APIManager = new APIManager();

    this.commandHandler = new CommandHandler(this, {
      directory: join(__dirname, '..', 'commands'),
      prefix: (msg: Message) => this.settings.get(msg.guild!, 'prefix', this.config.defaultPrefix),
      aliasReplacement: /-/g,
      allowMention: true,
      commandUtil: true,
      commandUtilLifetime: 3e5,
      defaultCooldown: 3000,
      argumentDefaults: {
        prompt: {
          modifyStart: async (msg: Message, text: string) =>
            await this.replyManager.modifyStart(msg, text),
          modifyRetry: async (msg: Message, text: string) =>
            await this.replyManager.modifyStart(msg, text),
          timeout: async (msg: Message) => await this.replyManager.timeout(msg),
          ended: async (msg: Message) => await this.replyManager.ended(msg),
          cancel: async (msg: Message) => await this.replyManager.cancel(msg),
          retries: 3,
          time: 20000
        },
        otherwise: ''
      }
    });

    this.inhibitorHandler = new InhibitorHandler(this, {
      directory: join(__dirname, '..', 'inhibitors')
    });

    this.listenerHandler = new ListenerHandler(this, {
      directory: join(__dirname, '..', 'listeners')
    });
  }

  private async init() {
    this.db = await Database.get('junko').connect();
    this.settings = new SettingsProvider(this.db.getRepository(Settings));
    this.APIManager.init();

    this.commandHandler.useInhibitorHandler(this.inhibitorHandler);
    this.commandHandler.useListenerHandler(this.listenerHandler);
    this.listenerHandler.setEmitters({
      commandHandler: this.commandHandler,
      inhibitorHandler: this.inhibitorHandler,
      listenerHandler: this.listenerHandler
    });

    this.commandHandler.loadAll();
    this.inhibitorHandler.loadAll();
    this.listenerHandler.loadAll();

    this.rpc.listen();
    this.prometheus.listen();
  }

  public async start() {
    await this.init();
    this.login(this.config.token);
  }
}
