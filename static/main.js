

const NotificationIcon = document.getElementById("icon")
const Notifications = document.getElementById("notifications")
const cancel_btn = document.getElementById("cancel_btn")

NotificationIcon.addEventListener("click", () => {
    Notifications.classList.add("show")
    NotificationIcon.classList.add("hidden")

})

cancel_btn.addEventListener("click", () => {
    Notifications.classList.remove("show")
    NotificationIcon.classList.remove("hidden")

})



