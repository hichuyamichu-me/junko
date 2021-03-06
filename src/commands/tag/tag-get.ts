import { Message } from 'discord.js';
import { Command } from 'discord-akairo';
import { Tag } from '../../models/Tag';

export default class TagGetCommand extends Command {
  public constructor() {
    super('tag-get', {
      category: 'tags',
      ownerOnly: false,
      channel: 'guild',
      description: {
        content: 'gets a tag',
        usage: '<tag>',
        examples: ['tagName']
      },
      args: [
        {
          id: 'name',
          type: 'lowercase',
          prompt: {
            start: 'Enter the tag name.',
            retry: 'You have to enter valid tag name.'
          }
        }
      ]
    });
  }

  public async exec(message: Message, { name }: { name: string }): Promise<void> {
    if (!name) return;
    const repo = this.client.db.getRepository(Tag);
    const tag = await repo
      .createQueryBuilder()
      .select('Tag.content', 'content')
      .where({ guild: message.guild.id, name })
      .getRawOne();

    if (tag) {
      message.util.send(tag.content);
    }
  }
}
