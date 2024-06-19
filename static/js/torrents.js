$(window).on('load', function () {
    updateInterface();
    changeTableView();
});


function bytesHumanReadable(bytes){

    if (!bytes){
        return '0 o/s';
    }
    const mega = 1024;

    if (Math.abs(bytes) < mega) {
        return `${bytes}o/s`;
    }

    const units = ['Ko', 'Mo', 'To'];
    let u = -1;
    let r = 10**1;
    do {
        bytes = bytes / mega;
        u = u + 1;
    }while (Math.round(Math.abs(bytes) * r) / r >= mega && u < units.length - 1);

    return `${bytes.toFixed(1)} ${units[u]}/s`;
}

$(window).resize(() => {
    changeTableView();
});

function changeTableView() {
    // if (window.innerWidth > 600) return false;
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

let toastTrigger = document.getElementById('liveToastBtn')
    let toastLiveExample = document.getElementById('liveToast')

function makeRequest(settings, callback){

    // get token from local storage
    let token = localStorage.getItem('token');
    if (token){
        settings.headers = {
            Authorization: 'Bearer ' + token
        }
    } else {
        // redirect to login page
        console.log("redirecting to login page")
        document.location.href = "/web/login";
    }

    $.ajax(settings).done(callback).fail((jqXHR, textStatus, errorThrown) => {

    console.log(jqXHR)
        let t = $("#toast-body");
        if ($(t).children().length === 0){
            let obj = jqXHR.responseText;
            let json = JSON.parse(obj);
            $(t).append(`<p>${json.error}</p>`);
            let toast = new bootstrap.Toast(toastLiveExample)

            toast.show()
        }


    });
}

function deleteTorrents(){
    let allCheckBox = $(".torrentCheckbox:checkbox:checked");
    for (let i = 0; i < allCheckBox.length; i++){
        let settings = {
            "url": `/api/v1/torrents/${$(allCheckBox[i]).attr('value')}`,
            "method": "DELETE",
            "timeout": 0,
        };

        makeRequest(settings, (response) => {
            $(allCheckBox[i]).parents("tr").remove();
        });

    }
}

function updateInterface(){
    let token = localStorage.getItem('token');
    var settings = {
        url: "/api/v1/torrents",
        method: "GET",
        timeout: 3000,
        headers: {
            Authorization: 'Bearer ' + token
        }
    };
    makeRequest(settings, (response, status, jqXHR) => {
        if (jqXHR.getResponseHeader('content-type').indexOf('text/html') >= 0){
            document.location.href = "/web/login";
        }
        var torrents = response;
        if ($("#tbodyTorrent tr").length === 0 || $(`#filename_${torrents[0].id}`).length === 0 ){
            $("#tbodyTorrent tr").remove();
            for (var i = 0; i < torrents.length; i ++){

                var torrent = torrents[i];
                var progress = torrent.progress || 100;
                var seeders = torrent.seeders || 0;
                var speed = bytesHumanReadable(torrent.speed) || 0;
                console.log(torrent.status);
                    let imgUri = encodeURI(`/static/img/${torrent.status}.png`);
                    $("#tableTorrent").append(`<tr>
                                        <td scope="row" id="filename_${torrent.id}"><p class="filenameTorrent">${torrent.filename}</p></td>
                                        <td><img src=""/><img width="24px" height="24px" id="status_${torrent.id}" src="${imgUri}" alt="${torrent.status}"/><span  id="progress_${torrent.id}">${progress}%</span></td>
                                        <td id="seeders_${torrent.id}">${seeders}</td>
                                        <td id="speed_${torrent.id}">${speed}</td>
                                        <td>
                                        <label for="download"><input class="torrentCheckbox" type="checkbox" value="${torrent.id}" name="id"/></label>
                                        </td>
                                    </tr>`);
            }
        }else {
            for (let i = 0; i < torrents.length; i++) {
                let torrent = torrents[i];
                let progress = $(`#progress_${torrent.id}`);
                let speed = $(`#speed_${torrent.id}`);
                let seeders = $(`#seeders_${torrent.id}`);
                let status = $(`#status_${torrent.id}`);

                status.attr("src", `/static/img/${torrent.status}.png`);
                progress.text(`${torrent.progress}%`);
                speed.text(bytesHumanReadable(torrent.speed) || 0);
                seeders.text(torrent.seeders || 0);
            }

        }
    });
}


setInterval (function () {
    updateInterface()
}, 2*1000);

function showLinks(links){
    console.log(links)
}

function showInfo(id){
    showModal();
}

function acceptAllFile(id){
    let settings = {
        url: `/api/v1/torrents/accept/${id}`,
        method: "GET",
        timeout: 0,
        processData: false,
    }

    makeRequest(settings, (result) => {
        console.log(result);
    });
}

function showModal(){
    this.blur();
    $.get(this.href, function(html){
        $(html).appendTo('body').modal();
    });
}

function closeModal(){
    modal.style.display = "none";
}

function uploadFile(){
    let form = new FormData();
    let files = $("#torrentFile")[0].files;
    for (let i = 0; i < files.length; i++){
        form.append("file", files[i], files[i].name);
    }

    var settings = {
        url: "/api/v1/torrent/upload",
        method: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        data: form,
    };

    $.ajax(settings).done(function (response) {
        let json = JSON.parse(response);
        for (let i = 0; i < json.length; i++){

            let settingsInfo = {
                url: `/api/v1/torrents/${json[i]['id']}`,
                method: "GET",
                timeout: 0,
            }

            makeRequest(settingsInfo, (r) => {
                acceptAllFile(r.id);
                $("#addingText").remove();
            });
        }
    });
}