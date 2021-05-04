const amqp = require("amqplib");
const msg = {number: 19}
connect();

async function connect() {
    try {
      const connection = await amqp.connect("amqp://localhost:5672");
      const channel = await connection.createChannel();
      const result = await channel.assertQueue("jobs");

      channel.consume("jobs", message => {
        const input = JSON.parse(message.content.toString());  
        console.log(`Received job with input ${input.number}`);
        if (input.number == 6)
          channel.ack(message); // ack means it's safe to remove msg from queue!
      })
      console.log("Client is waiting for messages...");
    }
    catch(ex){
        console.error(ex)
    }
}