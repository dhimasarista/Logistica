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

// Mengubah Integer to Rupiah
function formatRupiah(amount) {
    // Lakukan validasi untuk memastikan amount adalah tipe data numerik
    if (typeof amount !== 'number') {
        return 'Invalid input';
    }

    // Gunakan fungsi toLocaleString untuk memformat angka ke dalam format mata uang Rupiah
    return amount.toLocaleString('id-ID', {
        style: 'currency',
        currency: 'IDR',
        minimumFractionDigits: 0,
    });
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

function LoadingNotif(msg) {
    notyf.open({
        type: "info",
        background: "blue",   
        message: msg, 
    });
}

function clearCookies() {
    var cookies = document.cookie.split(";");

    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i];
        var eqPos = cookie.indexOf("=");
        var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT;path=/";
    }
}

// Contoh penggunaan
clearCookies();