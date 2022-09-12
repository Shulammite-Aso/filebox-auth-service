# Filebox Auth Service
This repo hosts the authentication service for **filebox**

### Meet the rest of the application:


- CLI tool --> [.../Shulammite-Aso/filebox](https://github.com/Shulammite-Aso/filebox)
- API gateway --> [.../Shulammite-Aso/filebox-api-gateway](https://github.com/Shulammite-Aso/filebox-api-gateway)
- Filebox service (the main "box" logic) --> [.../Shulammite-Aso/filebox-service](https://github.com/Shulammite-Aso/filebox-service)
- Email service --> [.../Shulammite-Aso/filebox-email-service](https://github.com/Shulammite-Aso/filebox-email-service)

## About Filebox 
Filebox is a cloud file sharing application hosted on kubernetes. It's an **original** project of mine, where i put to use some of the principles and technologies for 
building a truly cloud-native application. The rest of the readme will provide details about the architectural design, user actions, Flow, and technologies used.

## Architecture

Built on the microservices architecture and consist of three services, an api gateway and a CLI tool.  
  
  
**--diagram**


![filebox-diagram2](https://user-images.githubusercontent.com/48386390/187373095-7727a6cd-cd00-4b7a-bf46-58d3f4c7eb80.png)

## User Actions

Users can:

- Create an account which gives them access to a box
- Login to their account
- Send files to their own box
- Send files to the box of another user whose username they have  
- Delete a file from their box  
- Replace an existing file with new content
- Pull a file from their box
- Get a list of all files in their box (returns only file names)

Users cannot:

- Interact with the files of another user (can only send)
- Send a file to another users box if a file with such name already exists, they're asked to change the name

# Flow

The longest request journey happens when you send a file to the box of another user. It represents a complete journey through all the components, it goes this way:

- Login or register from the CLI tool to get authorization token
- Choose to send file to another user
- provide filepath and username of reciever
- POST request is made to the API gateway with authentication as bearer and request body containing, filename, file content as bytes, and  receiver username
- The API gateway calls the auth service to authorize the sending user using a middleware
- A second middleware verifies that the receiving user exists and retrieves their email
- The API gateway fowards the request to the filebox-service (containing the box logic), including the email that was retrieved into the request body in addition to 
the ones that came from the CLI tool
- The file is uploaded and the metadata stored in the receivers box
- Filebox-service makes request to the email service with receivers email, file name and name of sending user
- Email is sent to receiver informing them of the new file in their box and from whom it came.

## Technologies

- Go
- gRPC
- Postgresql
- MongoDB
- Docker
- Kubernetes

# Try It?

<a href="https://gitpod.io/#https://github.com/Shulammite-Aso/filebox"><img src="https://gitpod.io/button/open-in-gitpod.svg"/></a>

Sample files to start with can be found in the `sample_files` folder.

The services are running on gcp, if I'm not out of credit and everything works fine on my cluster, then the application would work just fine.
