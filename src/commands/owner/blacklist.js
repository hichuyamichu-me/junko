const { Command } = require('discord-akairo');

class BlacklistCommand extends Command {
  constructor() {
    super('blacklist', {
      aliases: ['blacklist'],
      ownerOnly: true,
      channel: ['guild', 'dm'],
      args: [
        {
          id: 'action',
          type: ['add', 'remove'],
          prompt: {
            start: message => `${message.author}, what are we doing this time?`,
            retry: message => `${message.author}, remember you can only add/remove from blacklist.`
          }
        },
        {
          id: 'user',
          type: 'user',
          prompt: {
            start: message => `${message.author}, who?`,
            retry: message => `${message.author}, seems like an invalid user.`
          }
        }
      ]
    });
  }

  async exec(message, { action, user }) {
    if (!user) return message.util.send('Please specify a valid user');
    switch (action) {
    case 'add':
      await this.client.store.hset('blacklist', user.id, 1);
      return message.util.send('That fucker is on my blacklist now!');
    case 'remove':
      await this.client.store.hdel('blacklist', user.id);
      return message.util.send('Removed from blacklist! I\'ll keep my an eye on him.');

    default:
      return message.util.send(`Unknown option:\`${action}\``);
    }
  }
}

module.exports = BlacklistCommand;
