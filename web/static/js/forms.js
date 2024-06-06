// Remove the text from submit buttons to show only the spinner during form requests
document.addEventListener('htmx:beforeRequest', function (e) {
  const t = getSubmitButtonText(e)
  t.style.display = 'none'
})

document.addEventListener('htmx:afterRequest', function (e) {
  const t = getSubmitButtonText(e)
  if (t) t.style.display = 'block'
})

function getSubmitButtonText(e) {
  const dispatcherEl = e.detail.elt
  if (dispatcherEl.tagName === 'FORM') {
    const submitButton = dispatcherEl.querySelector('button[type="submit"]')
    const submitButtonText = submitButton.querySelector('p')
    return submitButtonText
  }
}
