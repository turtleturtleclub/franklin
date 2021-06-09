const Discord = require('discord.js');
const { turtleid, prefix, token } = require('./config.json');
const client = new Discord.Client();

client.once('ready', () => {
        console.log('FRANKLIN IS HERE');
});

client.on('message', message => {
        for (let embed of message.embeds) {
                if (embed.title !== null && embed.title.includes('Free Tip appears')) {
                        message.channel.send('@here React now for some turtles! TURTLE,TURTLE! :turtle::turtle:');
                }
                else if(embed.description !== null && embed.description.includes('welcome to the TurtleTurtle.Club House')) {
                        message.channel.send('Thanks for coming new turtle. I am FRANKLIN the pond turtle!');
                }
        }
        if (message.content === '*test') {
                message.channel.send('OH YA TEST FRANKLIN IN DA POND');
                message.react('🐢');
        }
        else if (message.content === '*reg' || message.content === '*register') {
                message.channel.send(`Your Discord ID is ${message.author.id}, this is just a test and you are not registered`);
                message.react('🐢');
        }
        else if (message.content.startsWith(`${prefix}Block found`)) {
                message.channel.send('Nice job! :turtle::turtle: We found another!');
                message.react('🐢');
        }
        else if (message.content.startsWith(`${prefix}Block mined`)) {
                message.channel.send(':turtle::turtle: Block is good, turtles! :sunglasses: :turtle::turtle:');
                message.react('🐢');
        }
        else if (message.content.startsWith(`${prefix}Block orphaned`)) {
                message.channel.send(':cry: Next time, we got this.');
                message.react('😭');
        }
        else if (message.content === 'Hello Franklin' || message.content === 'hello Franklin' || message.content === 'Hey Franklin' || message.content === 'hey Franklin') {
                message.channel.send(`Turtle, turtle ${message.member.displayName}!`);
                message.react('🐢');
        }
        else if (message.content === 'Where is the block' || message.content === 'where is the block') {
                message.channel.send('Slow and steady wins the race.');
                message.react('🐢');
        }
        else if (message.content === 'come on Franklin' || message.content === 'Come on Franklin') {
                message.channel.send(`Slow and steady wins the race ${message.member.displayName}`);
                message.react('🐢');
        }
        else if (message.content === 'good boy Franklin' || message.content === 'Good boy Franklin') {
                message.channel.send(`Thank you ${message.member.displayName}, I try my turturtliest! :wink:`);
                message.react('🐢');
        }
        else if (message.content === 'franklin?' || message.content === 'FRANKLIN?' || message.content === 'Franklin?') {
                message.channel.send(`Do not worry, I am always here ${message.member.displayName} searching the pond for blocks.`);
               message.react('🐢');
        }
        else if (message.content === 'hugs' || message.content === 'Hugs' || message.content === 'hug' || message.content === 'Hug' || message.content === 'Hugs Franklin' || message.content === 'hug Franklin' || message.content === 'hug$                message.channel.send(`That was quite the turtley hug ${message.member.displayName}`);
                message.react('🐢');
        }
        else if (message.content === '🐢') {
                message.channel.send(`What are you doing out here? Did ${message.member.displayName} let you out? :eyes: Get back with your nest little one.`);
                message.react('🐢');
        }
        else if (message.content === '🐢🐢') {
                message.channel.send(`Looks to me ${message.member.displayName}, like you're turtley enough for the Turtle Turtle Club!`);
                message.react('🐢');
        }

});

client.login(token);
