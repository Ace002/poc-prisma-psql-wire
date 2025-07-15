# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Add this line to install OpenSSL for Prisma
RUN apk add --no-cache openssl

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install app dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Generate the Prisma client
RUN npx prisma generate

# Command to run the application
CMD ["npm", "start"]