import { Message } from 'discord.js';
import { Command, AkairoModule } from 'discord-akairo';
import { Logger } from '../../structs/Logger';

export default class ReloadCommand extends Command {
  public constructor() {
    super('reload', {
      aliases: ['reload', 'r'],
      category: 'owner',
      ownerOnly: true,
      description: {
        content: 'Reloads a module.',
        usage: '<module> [type:]',
        examples: ['eval', 'type:l cooldown']
      }
    });
  }

  public *args(): unknown {
    const type = yield {
      'match': 'option',
      'flag': ['--type='],
      'type': [['commandAlias', 'command', 'c'], ['inhibitor', 'i'], ['listener', 'l']],
      'default': 'command'
    };

    const mod = yield {
      type: (msg: Message, phrase: string) => {
        if (!phrase) return null;
        const resolver = this.handler.resolver.type(type);
        return resolver(msg, phrase);
      }
    };

    return { type, mod };
  }

  public exec(message: Message, { type, mod }: { type: string; mod: AkairoModule }): Promise<void> {
    if (!mod) {
      message.util.reply(
        `Invalid ${type} ${type === 'command' ? 'alias' : 'ID'} specified to reload.`
      );
      return;
    }

    try {
      mod.reload();
      message.util.reply(`Sucessfully reloaded ${type} \`${mod.id}\`.`);
      return;
    } catch (err) {
      Logger.error(err);
      message.util.reply(`Failed to reload ${type} \`${mod.id}\`.`);
      return;
    }
  }
}
