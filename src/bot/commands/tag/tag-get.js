const { Command } = require('discord-akairo');

class TagGetCommand extends Command {
  constructor() {
    super('tag-get', {
      category: 'tags',
      ownerOnly: false,
      channel: 'guild',
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

  async exec(message, { name }) {
    if (!name) return;
    const tag = await this.client.store.Tag.findOne({ where: { guildID: message.guild.id, name } });
    if (tag) {
      tag.toJSON();
      return message.util.send(tag.content);
    }
  }
}

module.exports = TagGetCommand;
