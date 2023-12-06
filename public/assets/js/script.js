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

function checkSession() {
    fetch("/check-session", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    })
    .then(response => response.json())
    .then(data => {
        if (data.sessionExists) {
            console.log("Sesi ada, pengguna dapat melanjutkan.");
        } else {
            return 0;
            Swal.fire({
                icon: "warning",
                title: "Session Ended",
                text: "Please, login again.",
                showCancelButton: false,
                confirmButtonColor: "#3085d6",
                confirmButtonText: "OK",
            }).then((result) => {
                if (result.isConfirmed || result.isDismissed) {
                    window.location.href = "/login";
                }
            });
        }
    })
    .catch(error => {
        console.error("Error:", error);
    });
}

document.addEventListener("DOMContentLoaded", function () {
    checkSession();
});


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