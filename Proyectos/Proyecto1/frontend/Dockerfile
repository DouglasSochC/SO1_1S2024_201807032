# Etapa de construcción
FROM node:18-alpine AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo de configuración de dependencias
COPY package*.json ./

# Instala las dependencias
RUN npm ci

# Copia el resto de los archivos de la aplicación
COPY . .

# Se crean las variables de entorno que estaran antes de la construccion de la aplicacion
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}

# Construye la aplicación para producción
RUN npm run build

# Etapa de producción
FROM nginx:1.21.4-alpine

# Puerto en el que la aplicación se ejecutará
EXPOSE 80

# Copia la configuración de Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copia los archivos construidos desde la etapa anterior
COPY --from=builder /app/out /usr/share/nginx/html

# Comando para iniciar la aplicación cuando se ejecute el contenedor
CMD ["nginx", "-g", "daemon off;"]