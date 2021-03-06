# Red Hat Bucket Store Service
This service provides a RESTful API to store, fetch and delete objects origanized in buckets.

## Contents
1. [API definition](#API-definition)
2. [Run the service locally](#run-the-service-locally)
3. [Try it out](#try-it-out)

<br/>

# API definition
## **Routes**
### /objects/{bucket}/{objectID}

#### **PUT**
##### Summary:

    Upload an object or replace it's content if it exists.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| bucket | path | Unique ID of the bucket | Yes | string |
| objectID | path | ID of the object to store | Yes | string |
| object | body |  | No | [object](#object) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Object created | [objectId](#objectId) |
| 500 | Failed to store object. |  |

#### **GET**
##### Summary:

    Download an object.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| bucket | path | Unique ID of the bucket | Yes | string |
| objectID | path | ID of the object to store | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [object](#object) |
| 404 | Object not found |  |
| 500 | Failed to fetch object. |  |

#### **DELETE**
##### Summary:

    Deletes an object.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| bucket | path | Unique ID of the bucket | Yes | string |
| objectID | path | ID of the object to store | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 404 | Object not found |
| 500 | Failed to delete object. |


## **Models**

### **object**

Object details

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| content | string | Text content of the object | No |

### **objectId**

Stored object id

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string | Stored object id | No | 

<br/>

# Run the service locally

Before starting, you can change the storage mode in *"resources/config/.object-store-service.yaml*", by setting the variable **STORAGE_TYPE** to either **pgsql** (to use postgresql storage) or **memory** (to use memory storage). 

You can also change the service port by changing the variable **SERVICE_PORT**.

  **1. Load tools**

Start by executing the following command to get the necessary tools:
    
    make tools

  **2. Run docker compose (Optional)**

Next, run this command to get PostgreSQL and PgAdmin up and running:

*Note: This is not necessary if you're running the service in memory mode*

    make run.infra.detach


  **3. Generate PostgreSQL golang models (Optional)**

Once everything is up and running, run this command to generate the golang database models:

*Note: This is only necessary if you change the sql models*

    make generate.models


<br/>

  **4. Get dependencies and build**

Next, get all dependencies needed and build the service executable using the following commands:

    make build.vendor
    make build.local

  **5. Run the service**

Now that everything is set, run the following command to launch the service:

    make run.local

<br/>

# Try it out

Now that everything is running, you can test the service using the following object:
```json
{
    "id": "42",
    "bucket": "bucket00",
    "content": "foo bar"
}
```

*Note: If you changed the service port, don't forget to override the **PORT** variable while running the commannds bellow.*

To start, you can run the following command to upload the object:

    make run.upload.object

Then to read the object content, you can use:

    make run.read.object

And finally, you can delete the object with this command:

    make run.delete.object
