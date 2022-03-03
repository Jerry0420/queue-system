const version: string = "/api/v1"

const generateURL = (route: string): string => {
    // const serverHost: string = (process.env.SERVER_HOST as string)
    let url = `${version}${route}`
    return url
}

const HTTPMETHOD = {
    "GET": "GET",
    "POST": "POST",
    "PUT": "PUT",
    "DELETE": "DELETE",
    "PATCH": "PATCH" 
}

const CONTENT_TYPE_JSON = {
    "Content-Type": "application/json"
}

const generateAuth = (token: string, withBearer: boolean=true) => {
    if (withBearer === true){
        token = `Bearer ${token}`
    }
    return {
        "Authorization": token
    }
}

interface RequestParams {
    endpoint: string
    method: string
    headers?: object
    body?: string
}

export {
    generateURL,
    HTTPMETHOD,
    CONTENT_TYPE_JSON,
    generateAuth,
    RequestParams
}