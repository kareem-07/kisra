console.log(`name: Bob
quantity: 15
date: 2025-06-01
paymentMethod: Zill
location: San Francisco
`) // demo values

const postOrder = document.getElementById("post-order")
if (postOrder === null) {
  throw new Error("Fatal: Cannot access 'post-order' element")
}

function postOrderOnSubmit(e: SubmitEvent) {
  e.preventDefault()

  const target = e.target
  if (target === null) {
    throw new Error("Event target not found")
  }

  if (!(target instanceof HTMLFormElement)) {
    throw new Error("Event target was expected to be an HTMLFormElement")
  }

  const formData = new FormData(target)

  const formDataObj = Object.fromEntries(formData.entries())
  Object.entries(formDataObj).forEach(([key, val]) => {
    if (val === null) {
      throw new Error(`The value of '${key}' is null`)
    }
    
    if (val === "") {
      throw new Error(`The value of '${key}' is an empty string`)
    }

    if (val instanceof File) {
      throw new Error(`The value of '${key}' is File`)
    }    

  })

  const formDataJSONStr = JSON.stringify(formDataObj)

  fetch(target.action, {
    method: target.method,
    headers: {
      "Content-Type": "application/json"
    },
    body: formDataJSONStr
  })
  .then(async res => {
    if (res.status >= 200 && res.status <= 299) {
      return res.text()
    }
    const err = await res.text()
    return await Promise.reject(err)
  })
  .then(console.log)
  .catch(err => {
    console.error({
      body: formDataJSONStr,
      error: err,
    })
  })

}

postOrder.addEventListener("submit", e => postOrderOnSubmit(e))