let abord = false;

function debridLink(link){
    let token = localStorage.getItem('token');
    let settings = {
        "url": "/api/v1/torrents/debrid",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json",
            "Authorization": "Bearer " +  token
        },
        "data": JSON.stringify({
            "link": link
        }),
    };
    let linkCheckbox = $("#show_only_link");
    let linkList = $("#link-list");
    $.ajax(settings).done(function (response) {
        if (!linkCheckbox.is(":checked")){
            linkList.append(
                generateDownloadCard(response.filename, response.filesize, response.download)
            )
        }else{
            linkList.append(`${response.download}\n`);
        }
    }).fail((jqXHR, other, other2) => {
        linkList.append(generateErrorCard(jqXHR.responseText));
    });
}

function generateErrorCard(link){
    return `<li class="list-group-item bg-dark text-white">
                    <div class="card" style="background-color: firebrick">
                      <div class="card-body">
                        <blockquote class="blockquote mb-0">
                            <p style="color: black;">${link}</p>
                        </blockquote>
                      </div>
                    </div>
                    </li>`
}


function generateDownloadCard(filename, fileSize, link){
    return `<li style="background-color: #a01414" class="list-group-item bg-dark text-white">
                    <div class="card">
                      <div class="card-body">
                        <blockquote class="blockquote mb-0">
                            <p style="color: black;">${filename} - (${bytesHumanReadable(fileSize)})</p>
                            <footer><a style="color: black;" href="${link}">${link}</a></footer>
                        </blockquote>
                      </div>
                    </div>
                    </li>`
}

function bytesHumanReadable(bytes){

    if (!bytes){
        return '0 o';
    }
    const mega = 1024;

    bytes /= 10;

    if (Math.abs(bytes) < mega) {
        return `${bytes}`;
    }

    const units = ['Ko', 'Mo', 'To'];
    let u = -1;
    let r = 10**1;
    do {
        bytes = bytes / mega;
        u = u + 1;
    }while (Math.round(Math.abs(bytes) * r) / r >= mega && u < units.length - 1);

    return `${bytes.toFixed(1)} ${units[u]}`;
}

function sleep(ms){
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function debridClick(event){
    resetDebrid();
    generateContainer();
    abord = false;
    let area = $("#links");

    if (area.val() === "" || area.val() === "\n"){
        return;
    }
    let links = area.val().split("\n");
    for (let i = 0; i < links.length; i++) {
        if (abord){
            return;
        }
        debridLink(links[i]);
        await sleep(1000);
    }
}


function generateContainer(){
    if ($("#link-list").length === 0){
        if ($("#show_only_link").is(":checked")){
            $("#list-card").append(`
                <div id="list-container" class="row align-items-center">
                    <div class="col-20">
                        <textarea style="height: 300px;" class="form-control" id="link-list"></textarea>
                    </div>
                </div>
                
            `);
        }else{
            $("#list-card").append(`
            <div id="list-container" class="row justify-content-center">
                <div class="col-6">
                    <div class="card bg-dark text-white"  style="width: 50rem;">
                      <ul class="bg-dark list-group list-group-flush" id="link-list">
            
                      </ul>
                    </div>
                </div>
            </div>
                
        `);
        }
    }
}

function resetDebrid(){
    const listContainer = $("#list-container");
    abord = true;
    if (listContainer.length !== 0) {
        listContainer.remove();
    }
}

