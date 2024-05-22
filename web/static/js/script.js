// Remove the text from submit buttons to show only the spinner during form requests
document.addEventListener("htmx:beforeRequest", function () {
  const button = document.querySelector('button[type="submit"]')
  button.childNodes.forEach((child) => {
    if (child.nodeType === Node.TEXT_NODE && child.textContent.trim() !== "") {
      child.textContent = "" // Remove the text
    }
  })
})
