import * as httpTools from './base'

interface Customer {
    name: string
    phone: string
    queue_id: number
  }

const createCustomers = (sessionId: string, storeId: number, customers: Customer[]): Promise<any> => {
    const jsonBody: string = JSON.stringify({
        "store_id": storeId,
        "customers": customers
    })
    return fetch(
        httpTools.generateURL("/customers"), { 
            method: httpTools.HTTPMETHOD.POST,
            headers: {...httpTools.CONTENT_TYPE_JSON, ...httpTools.generateAuth(sessionId, false)},
            body: jsonBody
        }
    )
      .then(response => response.json())
      .then(jsonResponse => {
          console.log(jsonResponse)
          return jsonResponse
      })
      .catch(error => {
          console.error(error)
          throw new Error("createCustomers error")  
      })
}

const updateCustomer = (customerId: number, normalToken: string, storeId: number, queueId: number, oldCustomerStatus: string, newCustomerStatus: string): Promise<any> => {
    const route = "/customers/".concat(customerId.toString())
    const jsonBody: string = JSON.stringify({
        "store_id": storeId,
        "queue_id": queueId,
        "old_customer_status": oldCustomerStatus,
        "new_customer_status": newCustomerStatus,
    })
    return fetch(
        httpTools.generateURL(route), { 
            method: httpTools.HTTPMETHOD.PUT,
            headers: {...httpTools.CONTENT_TYPE_JSON, ...httpTools.generateAuth(normalToken)},
            body: jsonBody
        }
    )
      .then(response => response.json())
      .then(jsonResponse => {
          console.log(jsonResponse)
          return jsonResponse
      })
      .catch(error => {
          console.error(error)
          throw new Error("updateCustomer error")  
      })
}

export {
    createCustomers,
    updateCustomer
}