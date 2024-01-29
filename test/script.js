document.getElementById('uploadForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Предотвращаем стандартное поведение отправки формы

    var fileInput = document.getElementById('fileInput');
    var files = fileInput.files;

    if (files.length === 0) {
        alert('Please select at least one image to upload.');
        return;
    }

    var formData = new FormData();

    for (var i = 0; i < files.length; i++) {
        var file = files[i];
        formData.append('images[]', file);
    }

    fetch('http://localhost:8081/api/images', {
        method: 'POST',
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            console.log('Upload successful:', data);
            document.getElementById('status').innerText = 'Upload successful';
        })
        .catch(error => {
            console.error('Error uploading:', error);
            document.getElementById('status').innerText = 'Error uploading';
        });
});
