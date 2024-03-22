const Redis = require('ioredis');
const conexion = new Redis({
    host: '10.35.13.51',
    port: 6379,
    connectTimeout: 5000,
});

function hacerPub() {
    conexion.publish('test', 'Hola a todos')
        .then(() => {
            console.log('Mensaje enviado');
        })
        .catch((error) => {
            console.error('Error al enviar mensaje', error);
        });
}

setInterval(hacerPub, 3000);