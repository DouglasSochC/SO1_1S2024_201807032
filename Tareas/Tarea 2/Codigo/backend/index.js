const express = require('express');
const mongoose = require('mongoose');
const app = express();
const cors = require('cors');
const PORT = 3001;

// Configurar CORS
app.use(cors());
app.use(express.json({ limit: '5mb' }));
app.use(express.urlencoded({ limit: '5mb' }));

// Conexión a la base de datos MongoDB
mongoose.connect('mongodb://MongoDB:27017/tarea2', {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

// Definir el modelo de la colección
const imagenSchema = new mongoose.Schema({
  nombre: { type: String, default: 'Archivo' + Date.now() },
  imagen: String,
  fechaCarga: { type: Date, default: Date.now },
});

const Imagen = mongoose.model('Imagen', imagenSchema, 'imagenes');

// Ruta para almacenar una imagen
app.post('/subir-imagen', async (req, res) => {
  try {
    const { imagenBase64 } = req.body;

    // Crear un documento de imagen
    const nuevaImagen = new Imagen({
      imagen: imagenBase64,
    });

    // Guardar la informacion
    await nuevaImagen.save();

    res.json({ mensaje: 'Imagen almacenada exitosamente' });
  } catch (error) {
    console.error(error);
    res.status(500).json({ mensaje: 'Error al almacenar la imagen' });
  }
});

// Ruta para obtener todos los registros almacenados
app.get('/obtener-registros', async (req, res) => {
  try {
    // Obtener todos los documentos de la colección
    const registros = await Imagen.find();
    res.json(registros);
  } catch (error) {
    console.error(error);
    res.status(500).json({ mensaje: 'Error al obtener los registros' });
  }
});

// Iniciar el servidor
app.listen(PORT, () => {
  console.log(`Servidor escuchando en el puerto ${PORT}`);
});