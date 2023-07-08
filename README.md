# Message-Queuing-System
Approach
Flow Diagram:
![image](https://github.com/mkshailesh100/Message-Queuing-System/assets/59009436/30e67527-a518-4357-bdc6-8b4624a2a5f4)


Explanation:
We have created an API which will take following parameters as input- user_id, product_name, product_description - product_images, product_price and store them in product table. 
Then producer.go is run which takes the product id and populate it into RabbitMQ queue. Later consumer.go runs and fetch the product id from queue and the respective record from the database.
The fetch URL is passed to compressor which download the images in Images folder, do the compression and store them in CompressedImage folder and then insert the images path in the database as required.

Below are the two table created in database named product_management
Tables: Users and Products


Below is the project structure and file explaination
- cmd/
  - main.go   Starting point of Project (Run go run cmd/main.go to run this project)
- internal/
  - api/
    - handlers.go    		API function 
  - db/
    - database.go   		Database creating, migration and other functions 
  - messaging/
    - producer.go		To send the productId to rabbitMq  
    - consumer.go		To fetch the product from rabbitMq
  - compression/
    - compressor.go		To download image and store the location in database
- compressedImages
  - All images
- images
  All images
- pkg/
  - models/
    - user.go			Structure of user table
    - product.go		Structure of product table
- tests/
  - integration/
    - integration_test.go	Unit_testing file
  - benchmark/
    - benchmark_test.go	Benchmark_testing file

CURL request to hit
curl --location --request POST 'http://localhost:8080/products' \
--header 'Content-Type: application/json' \
--data-raw '{
   "user_id": 1,
   "product_name": "Sample Product",
   "product_description": "This is a sample product",
   "product_images": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSB-Sl-zoom2cnZQvm8yKk409Pg0ts0gZ7pxEAvOL38oQ&s,https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTNtCh17cCUl3OeiiqnqYb72OPfHLLRVte3sg5Lz5duGg&s",
   "product_price": 9.99
}'
Note: Product Images is passed as string separated URL which is then internally converted to Array and processed for consumer to download and compressed the imahe

Command:
To run project: go run cmd/main.go
To run unit_test: run integration_test.go


