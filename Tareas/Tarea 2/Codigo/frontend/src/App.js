import React, { useState, useEffect } from 'react';
import Webcam from 'react-webcam';
import axios from 'axios';
import './App.css';

const App = () => {
  const [image, setImage] = useState(null);
  const [uploadMessage, setUploadMessage] = useState(null);
  const [registros, setRegistros] = useState([]);

  const webcamRef = React.useRef(null);

  const capture = React.useCallback(() => {
    const imageSrc = webcamRef.current.getScreenshot();
    setImage(imageSrc);
  }, [webcamRef]);

  const uploadImage = async () => {
    try {
      if (image) {
        const formData = {
          imagenBase64: image,
        };

        const response = await axios.post('http://localhost:3001/subir-imagen', formData);
        const serverMessage = response.data.mensaje;

        setUploadMessage(serverMessage);
        setImage(null);
        // Recargar los registros después de subir una nueva imagen
        obtenerRegistros();
      } else {
        console.error('Captura una imagen antes de intentar subirla');
      }
    } catch (error) {
      console.error('Error al enviar la imagen', error);
      setUploadMessage('Error al enviar la imagen');
    }
  };

  const obtenerRegistros = async () => {
    try {
      const response = await axios.get('http://localhost:3001/obtener-registros');
      setRegistros(response.data);
    } catch (error) {
      console.error('Error al obtener los registros', error);
    }
  };

  useEffect(() => {
    // Obtener registros cuando el componente se monta
    obtenerRegistros();
  }, []); // La dependencia vacía asegura que se ejecute solo una vez al montar el componente

  return (
    <div className="App">
      <b>Nombre:</b> Douglas Alexander Soch Catalán
      <br />
      <b>Carnet:</b> 201807032
      <div className="WebcamCapture">
        <Webcam
          audio={false}
          ref={webcamRef}
          screenshotFormat="image/png"
        />
        <button onClick={capture}>Capturar Foto</button>
      </div>
      {
        image ? (
          <div className="ImagePreview">
            <h2>Vista Previa</h2>
            <img src={image} alt="Captured" />
            <br />
            <button onClick={uploadImage}>Subir Imagen</button>
          </div>
        ) : (
          uploadMessage && (
            <div className="UploadMessage">
              <h2>{uploadMessage}</h2>
            </div>)
        )
      }

      <div>
        {registros.length > 0 && (
          <div className="Registros">
            <h2>Registros</h2>
            <table>
              <thead>
                <tr>
                  <th>Nombre</th>
                  <th>Fecha de Carga</th>
                  <th>Imagen</th>
                </tr>
              </thead>
              <tbody>
                {registros.map((registro) => (
                  <tr key={registro._id}>
                    <td>{registro.nombre}</td>
                    <td>{new Date(registro.fechaCarga).toLocaleString()}</td>
                    <td>
                      <img
                        src={registro.imagen}
                        alt={registro.nombre}
                        style={{ maxWidth: '100px', maxHeight: '100px' }}
                      />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
};

export default App;