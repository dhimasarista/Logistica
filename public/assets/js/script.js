function checkFileSize(file) {
    const maxSize = 200 * 1024;
    if (file.files.length > 0) {
        const fileSize = file.files[0].size;
        if (fileSize > maxSize) {
            alert("File Tidak Boleh Lebih Dari 200kb");

            file.value = "";
        }
    }                                   
}

function checkFileType(file, expectedTypes, errorMsg, toDelete) {
    if (!expectedTypes.includes(file.type)) {
        alert(errorMsg);
        toDelete.value = null;
        return;
    }
}