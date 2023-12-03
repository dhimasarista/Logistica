function checkFileSize(file) {
    const maxSize = 200 * 1024;
    if (file.files.length > 0) {
        const fileSize = file.files[0].size;
        if (fileSize > maxSize) {
            alert("File Tidak Boleh Lebih Dari 200kb");

            file.value = "";
            return 0;
        }
    }                                   
}

function checkFileType(file, expectedTypes, errorMsg, toDelete) {
    if (!expectedTypes.includes(file.type)) {
        alert(errorMsg);
        toDelete.value = null;
        return 0;
    }
}

// let sidebarIsToggled = localStorage.getItem("toggle-mode") === "true";
// const sidebarClass = document.querySelector(".sidebar");

// function sidebarToggledEvent() {
//     if (!sidebarIsToggled) {
//         localStorage.setItem("toggle-mode", "true");
//         sidebarIsToggled = true;
//         sidebarClass.classList.add("toggled");
//     } else {
//         localStorage.setItem("toggle-mode", "false");
//         sidebarIsToggled = false;
//         sidebarClass.classList.remove("toggled");
//     }
// }

// document.addEventListener("DOMContentLoaded", () => {
//     const isToggled = localStorage.getItem("toggle-mode")
//     if (isToggled) {
//         sidebarClass.classList.add("toggled")
//     } else if (!isToggled) {
//         sidebarClass.classList.remove("toggled")
//     }
// })