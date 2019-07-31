const { Command } = require('discord-akairo');

class DelTagCommand extends Command {
  constructor() {
    super('tag-del', {
      ownerOnly: false,
      channel: 'guild',
      args: [
        {
          id: 'name',
          type: 'string',
          prompt: {
            start: message => `${message.author}, enter the tag name.`,
            retry: message => `${message.author}, you have to enter valid tag name.`
          }
        }
      ]
    });
  }

  async exec(message, { name }) {
    const tag = await this.client.store.hdelAsync(message.guild.id, `tag-${name}`);
    if (tag) {
      return message.util.send(`Succesfuly deleted \`${name}\` `);
    }
    return message.util.send('No such tag');
  }
}

module.exports = DelTagCommand;
