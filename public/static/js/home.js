let abord = false;

function debridLink(link) {
    const token = localStorage.getItem('token');
    console.log(token);

    const settings = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + token
        },
        body: JSON.stringify({ link: link })
    };

    const linkCheckbox = document.getElementById('show_only_link');
    const linkList = document.getElementById('link-list');

    fetch("/api/v1/torrents/debrid", settings)
        .then(response => response.json())
        .then(data => {
            if (!linkCheckbox.checked) {
                linkList.innerHTML += generateDownloadCard(data.filename, data.fileSize, data.download);
            } else {
                linkList.innerHTML += `${data.download}\n`;
            }
        })
        .catch(error => {
            linkList.innerHTML += generateErrorCard(error.message);
        });
}

function generateErrorCard(link) {
    return `<li class="list-group-item bg-dark text-white">
                <div class="card" style="background-color: firebrick">
                  <div class="card-body">
                    <blockquote class="blockquote mb-0">
                        <p style="color: black;">${link}</p>
                    </blockquote>
                  </div>
                </div>
              </li>`;
}

function generateDownloadCard(filename, fileSize, link) {
    return `<li style="background-color: #a01414" class="list-group-item bg-dark text-white">
                <div class="card">
                  <div class="card-body">
                    <blockquote class="blockquote mb-0">
                        <p style="color: black;">${filename} - (${bytesHumanReadable(fileSize)})</p>
                        <footer><a style="color: black;" href="${link}">${link}</a></footer>
                    </blockquote>
                  </div>
                </div>
              </li>`;
}

function bytesHumanReadable(bytes) {
    if (!bytes) {
        return '0 o';
    }
    const mega = 1024;
    bytes /= 10;

    if (Math.abs(bytes) < mega) {
        return `${bytes}`;
    }

    const units = ['Ko', 'Mo', 'To'];
    let u = -1;
    let r = 10 ** 1;
    do {
        bytes = bytes / mega;
        u = u + 1;
    } while (Math.round(Math.abs(bytes) * r) / r >= mega && u < units.length - 1);

    return `${bytes.toFixed(1)} ${units[u]}`;
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function debridClick(event) {
    resetDebrid();
    generateContainer();
    abord = false;

    const area = document.getElementById('links');
    if (!area.value || area.value === "\n") {
        return;
    }

    const links = area.value.split("\n");
    for (let i = 0; i < links.length; i++) {
        if (abord) {
            return;
        }
        debridLink(links[i]);
        await sleep(1000);
    }
}

function generateContainer() {
    const listCard = document.getElementById('list-card');
    if (!document.getElementById('link-list')) {
        if (document.getElementById('show_only_link').checked) {
            listCard.innerHTML += `
                <div id="list-container" class="row align-items-center">
                    <div class="col-20">
                        <textarea style="height: 300px;" class="form-control" id="link-list"></textarea>
                    </div>
                </div>
            `;
        } else {
            listCard.innerHTML += `
            <div id="list-container" class="row justify-content-center">
                <div class="col-6">
                    <div class="card bg-dark text-white" style="width: 50rem;">
                      <ul class="bg-dark list-group list-group-flush" id="link-list"></ul>
                    </div>
                </div>
            </div>
        `;
        }
    }
}

function resetDebrid() {
    const listContainer = document.getElementById('list-container');
    abord = true;
    if (listContainer) {
        listContainer.remove();
    }
}
