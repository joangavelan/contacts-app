// Hide text elements on HTMX request triggers to show only the spinner during requests
document.addEventListener('htmx:beforeRequest', function (e) {
  const elt = getTextElementFromDispatcher(e.detail.elt)
  if (elt) elt.style.display = 'none'
})

document.addEventListener('htmx:afterRequest', function (e) {
  const elt = getTextElementFromDispatcher(e.detail.elt)
  if (elt) elt.style.display = 'block'
})

function getTextElementFromDispatcher(dispatcherEl) {
  if (dispatcherEl.tagName === 'FORM') {
    const submitButton = dispatcherEl.querySelector('button[type="submit"]')
    const submitButtonTextEl = submitButton.querySelector('p')
    return submitButtonTextEl
  }

  return dispatcherEl.querySelector('p')
}
