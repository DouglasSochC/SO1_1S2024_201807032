import redis # pip install redis

# Crear una conexión a Redis
conexion = redis.StrictRedis(host='10.35.13.51', port=6379, db=0)

# Suscribirse al canal 'test'
pubsub = conexion.pubsub()
pubsub.subscribe('test')

# Manejar los mensajes recibidos
for mensaje in pubsub.listen():
    if mensaje['type'] == 'message':
        print(f"Recibido mensaje de {mensaje['channel'].decode('utf-8')}: {mensaje['data'].decode('utf-8')}")

# Manejar errores
def manejar_error(err):
    print(f"Error en la conexión: {err}")
pubsub.on_error(callback=manejar_error)