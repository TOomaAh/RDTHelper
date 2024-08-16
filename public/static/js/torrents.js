const imageMap = new Map();

window.addEventListener('load', () => {
    updateInterface();
    changeTableView();
});

window.addEventListener('resize', changeTableView);

function bytesHumanReadable(bytes) {
    if (!bytes) {
        return '0 o/s';
    }
    const mega = 1024;

    if (Math.abs(bytes) < mega) {
        return `${bytes}o/s`;
    }

    const units = ['Ko', 'Mo', 'To'];
    let u = -1;
    let r = 10 ** 1;
    do {
        bytes = bytes / mega;
        u = u + 1;
    } while (Math.round(Math.abs(bytes) * r) / r >= mega && u < units.length - 1);

    return `${bytes.toFixed(1)} ${units[u]}/s`;
}

function changeTableView() {
    const tableEl = document.querySelector("table");
    const thEls = tableEl.querySelectorAll('thead th');
    const tdLabels = Array.from(thEls).map(el => el.innerText);
    tableEl.querySelectorAll('tbody tr').forEach(tr => {
        Array.from(tr.children).forEach(
            (td, ndx) => {
                td.setAttribute('label', tdLabels[ndx]);
            }
        );
    });
}

function makeRequest(settings, callback) {
    let token = localStorage.getItem('token');
    if (token) {
        settings.headers = {
            Authorization: 'Bearer ' + token
        };
    } else {
        console.log("redirecting to login page");
        window.location.href = "/login";
        return;
    }

    fetch(settings.url, {
        method: settings.method,
        headers: settings.headers,
        body: settings.data
    })
    .then(response => {
        if (response.headers.get('content-type').indexOf('text/html') >= 0) {
            window.location.href = "/login";
            return;
        }
        return response.json();
    })
    .then(data => callback(data))
    .catch(error => {
        console.log(error);
        let toastBody = document.getElementById('toast-body');
        if (toastBody.children.length === 0) {
            let json = JSON.parse(error.responseText);
            toastBody.innerHTML = `<p>${json.error}</p>`;
            let toast = new bootstrap.Toast(document.getElementById('liveToast'));
            toast.show();
        }
    });
}

function deleteTorrents() {
    let allCheckBox = document.querySelectorAll(".torrentCheckbox:checked");
    allCheckBox.forEach(checkbox => {
        let settings = {
            url: `/api/v1/torrents/${checkbox.value}`,
            method: "DELETE",
            headers: {
                'Content-Type': 'application/json'
            }
        };

        makeRequest(settings, (response) => {
            checkbox.closest("tr").remove();
        });
    });
}

function updateInterface() {
    let token = localStorage.getItem('token');
    var settings = {
        url: "/api/v1/torrents",
        method: "GET",
        headers: {
            Authorization: 'Bearer ' + token
        }
    };

    makeRequest(settings, (response) => {
        var torrents = response;
        var tbodyTorrent = document.getElementById("tbodyTorrent");
        var tableTorrent = document.getElementById("tableTorrent");

        if (tbodyTorrent.querySelectorAll("tr").length === 0 || !document.getElementById(`filename_${torrents[0].id}`)) {
            tbodyTorrent.innerHTML = '';
            var fragment = document.createDocumentFragment();

            torrents.forEach(torrent => {
                var progress = torrent.progress || 100;
                var seeders = torrent.seeders || 0;
                var speed = bytesHumanReadable(torrent.speed) || 0;
                let imgUri = encodeURI(`/static/img/${torrent.status}.png`);

                var row = document.createElement("tr");
                row.innerHTML = `
                    <td scope="row" id="filename_${torrent.id}"><p class="filenameTorrent">${torrent.filename}</p></td>
                    <td><img src="" /><img width="24px" height="24px" id="status_${torrent.id}" src="${imgUri}" alt="${torrent.status}" /><span id="progress_${torrent.id}">${progress}%</span></td>
                    <td id="seeders_${torrent.id}">${seeders}</td>
                    <td id="speed_${torrent.id}">${speed}</td>
                    `
                    if (torrent.status === "downloaded") {
                    row.innerHTML += `<td>
                            <label for="download"><input class="torrentCheckbox" type="checkbox" value="${torrent.id};${torrent.links}" name="id" /></label>
                        </td>`
                    }
                ;
                fragment.appendChild(row);
            });

            tableTorrent.appendChild(fragment);
        } else {
            torrents.forEach(torrent => {
                let progress = document.getElementById(`progress_${torrent.id}`);
                let speed = document.getElementById(`speed_${torrent.id}`);
                let seeders = document.getElementById(`seeders_${torrent.id}`);
                let status = document.getElementById(`status_${torrent.id}`);

                status.setAttribute("src", `/static/img/${torrent.status}.png`);
                progress.textContent = `${torrent.progress}%`;
                speed.textContent = bytesHumanReadable(torrent.speed) || 0;
                seeders.textContent = torrent.seeders || 0;
            });
        }
    });
}

setInterval(updateInterface, 2 * 1000);

function showLinks(links) {
    console.log(links);
}

function showInfo(id) {
    showModal();
}

function acceptAllFile(id) {
    let settings = {
        url: `/api/v1/torrents/accept/${id}`,
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    };

    makeRequest(settings, (result) => {
        console.log(result);
    });
}

function showModal() {
    this.blur();
    fetch(this.href)
        .then(response => response.text())
        .then(html => {
            const modal = document.createElement('div');
            modal.innerHTML = html;
            document.body.appendChild(modal);
            const bootstrapModal = new bootstrap.Modal(modal.querySelector('.modal'));
            bootstrapModal.show();
        });
}

function closeModal() {
    modal.style.display = "none";
}

function uploadFile() {
    let token = localStorage.getItem('token');
    if (!token) {
        console.log("redirecting to login page");
        window.location.href = "/login";
        return;
    }

    let form = new FormData();
    let files = document.getElementById("torrentFile").files;
    console.log(files.length);
    for (let i = 0; i < files.length; i++) {
        form.append("file", files[i], files[i].name);
    }

    console.log(form);

    var settings = {
        url: "/api/v1/torrent/upload",
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + token,
        },
        body: form,
    };


    fetch(settings.url, {
        method: settings.method,
        headers: settings.headers,
        body: settings.body
    })
    .then(response => response.json());
}
