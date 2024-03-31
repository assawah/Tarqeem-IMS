

const Notifications = document.getElementById("icon")
const NotificationIcon = document.getElementById("notifications")

NotificationIcon.addEventListener("click", () => {
    console.log("clicked")
    Notifications.classList.add("show")
    NotificationIcon.classList.add("hidden")

})

