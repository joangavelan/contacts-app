const toastContainer = document.querySelector('#toast-container')

class Toast {
  /**
   * A class representing a Toast notification.
   * @param variant {("info"|"success"|"warning"|"error")}
   * @param message { string }
   */
  constructor(variant, message) {
    this.variant = variant
    this.message = message
  }

  /**
   * Makes the toast element. A span element containing the entire notification.
   * @returns {HTMLSpanElement}
   */
  #makeToastElement() {
    const span = document.createElement('span')
    span.className = `toast toast-${this.variant}`
    span.textContent = this.message
    return span
  }

  /**
   * Displays the toast notification to the user.
   */
  show() {
    const toastEl = this.#makeToastElement()
    toastContainer.appendChild(toastEl)
  }
}

/**
 * Listen for the custom 'triggerToast' event.
 *
 * The 'triggerToast' event is triggered by an htmx response header, specifically the "HX-Trigger" response header.
 * When this event is detected, a new Toast object is created using the event details (variant and message).
 * The toast notification is then displayed using the toast.show() method.
 */
document.addEventListener('triggerToast', (e) => {
  const toast = new Toast(e.detail.variant, e.detail.message)
  toast.show()
})

/**
 * Handle toast notifications with animations and automatic removal.
 *
 * Constants:
 * - TOAST_DISPLAY_TIME: Total display time (in milliseconds).
 * - TOAST_TRANSITION_TIME: Transition time (in milliseconds).
 *
 * The MutationObserver observes changes in the toast container:
 * - Triggers 'slide-in' animation shortly after a toast is added.
 * - Triggers 'fade-out' animation before the total display time elapses.
 * - Removes the toast from the DOM after the total display time.
 */

const TOAST_DISPLAY_TIME = 2500
const TOAST_TRANSITION_TIME = 500

const observer = new MutationObserver((mutationList) => {
  const addedToast = mutationList[0].addedNodes.item(0)

  if (addedToast) {
    setTimeout(() => {
      addedToast.classList.add('slide-in')
    }, 50)

    setTimeout(() => {
      addedToast.classList.add('fade-out')
    }, TOAST_DISPLAY_TIME - TOAST_TRANSITION_TIME)

    setTimeout(() => {
      addedToast.remove()
    }, TOAST_DISPLAY_TIME)
  }
})

observer.observe(toastContainer, { childList: true })
