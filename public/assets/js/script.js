// Memeriksa ukuran file yang diunggah
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
// Memeriksa jenis tipe file yang akan diunggah
function checkFileType(file, expectedTypes, errorMsg, toDelete) {
    if (!expectedTypes.includes(file.type)) {
        alert(errorMsg);
        toDelete.value = null;
        return 0;
    }
}
// Pustaka notifikasi
let notyf = new Notyf({
    duration: 4000,
    dismissible: true,
    position: {
        x: "right",
        y: "top"
    }
});

function ErrorNotif(error) {
    notyf.open({
        type: "error",
        background: "orange",   
        message: error, 
    });
}

function InternalServerError(error) {
    notyf.open({
        type: "error",
        message: error,
    });
}
function ServerStatusOke(msg) {
    notyf.open({
        type: "success",
        message: msg,
    })
}