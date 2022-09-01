// GENERATED BY textFileToGoConst
// GitHub:     github.com/logrusorgru/textFileToGoConst
// input file: html\front\website.js
// generated:  Wed Aug 24 15:02:56 CEST 2022

package website

const JS = "let gWallets = [];\r\nlet deployers = [];\r\nlet actualTxType = '';\r\nlet nextFileToUpload;\r\n\r\n// INIT\r\ngetWallets();\r\ngetWebsiteDeployerSC();\r\ninitializeDefaultWallet();\r\nsetupModal();\r\n\r\n// Write the default wallet text in wallet popover component\r\nasync function getWebsiteDeployerSC() {\r\n\tlet defaultWallet = getDefaultWallet();\r\n\r\n\t$('#website-deployers-table tbody tr').remove();\r\n\r\n\taxios\r\n\t\t.get('/my/domains/' + defaultWallet)\r\n\t\t.then((websites) => {\r\n\t\t\tlet count = 0;\r\n\t\t\tfor (const website of websites.data) {\r\n\t\t\t\ttableInsert(website, count);\r\n\t\t\t\tcount++;\r\n\t\t\t}\r\n\t\t\tdeployers = websites.data;\r\n\t\t})\r\n\t\t.catch((e) => {\r\n\t\t\terrorAlert(e.response.data.code);\r\n\t\t});\r\n}\r\n\r\n// Write the default wallet text in wallet popover component\r\nfunction initializeDefaultWallet() {\r\n\tlet defaultWallet = getDefaultWallet();\r\n\tif (defaultWallet === '') {\r\n\t\tdefaultWallet = 'Connect';\r\n\t}\r\n\t$('.popover__title').html(defaultWallet);\r\n}\r\n\r\n// Retrieve the default wallet nickname in cookies\r\nfunction getDefaultWallet() {\r\n\tlet defaultWallet = '';\r\n\tconst cookies = document.cookie.split(';');\r\n\tcookies.forEach((cookie) => {\r\n\t\tconst keyValue = cookie.split('=');\r\n\t\tif (keyValue[0] === 'defaultWallet') {\r\n\t\t\tdefaultWallet = keyValue[1];\r\n\t\t}\r\n\t});\r\n\treturn defaultWallet;\r\n}\r\n\r\nfunction setupModal() {\r\n\t$('#passwordModal').on('shown.bs.modal', function () {\r\n\t\t$('#passwordModal').trigger('focus');\r\n\t});\r\n}\r\n\r\nfunction setTxType(txType) {\r\n\tactualTxType = txType;\r\n}\r\n\r\nasync function callTx() {\r\n\tconst passwordValue = $('#walletPassword').val();\r\n\r\n\tif (actualTxType === 'deployWebsiteCreator') {\r\n\t\tdeployWebsiteDeployerSC(passwordValue);\r\n\t}\r\n\tif (actualTxType.includes('uploadWebsiteCreator')) {\r\n\t\tconst websiteIndex = actualTxType.split('uploadWebsiteCreator')[1];\r\n\t\tuploadWebsite(nextFileToUpload, websiteIndex, passwordValue);\r\n\t}\r\n}\r\n\r\n// open file upload\r\nfunction openDialog() {\r\n\tdocument.getElementById('fileid0').value = null;\r\n\tdocument.getElementById('fileid0').click();\r\n}\r\n\r\n// Handle event on file selecting\r\nfunction handleFileSelect(evt) {\r\n\tlet files = evt.target.files; // get files\r\n\tlet f = files[0];\r\n\tconst reader = new FileReader();\r\n\treader.onload = (event) => importWallet(JSON.parse(event.target.result)); // desired file content\r\n\treader.onerror = (error) => reject(error);\r\n\treader.readAsText(f);\r\n}\r\n\r\nfunction errorAlert(error) {\r\n\tdocument.getElementsByClassName('alert-danger')[0].style.display = 'block';\r\n\tdocument.getElementsByClassName('alert-danger')[0].innerHTML = error;\r\n\tsetTimeout(function () {\r\n\t\tdocument.getElementsByClassName('alert-danger')[0].style.display = 'none';\r\n\t}, 5000);\r\n}\r\n\r\nfunction successMessage(message) {\r\n\tdocument.getElementsByClassName('alert-primary')[0].style.display = 'block';\r\n\tdocument.getElementsByClassName('alert-primary')[0].innerHTML = message;\r\n\tsetTimeout(function () {\r\n\t\tdocument.getElementsByClassName('alert-primary')[0].style.display = 'none';\r\n\t}, 5000);\r\n}\r\n\r\n// Append wallet accounts in popover component list\r\nasync function feedWallet(w) {\r\n\tlet counter = 0;\r\n\tfor (const wallet of w) {\r\n\t\t$('#wallet-list').append(\r\n\t\t\t\"<li class='wallet-item'><a class='wallet-link' id='wallet-link-\" +\r\n\t\t\t\tcounter +\r\n\t\t\t\t\"' onclick='changeDefaultWallet(event)' href='#'>\" +\r\n\t\t\t\twallet.nickname +\r\n\t\t\t\t'</a></li>'\r\n\t\t);\r\n\t\tcounter++;\r\n\t}\r\n}\r\n\r\n// Handle popover click & update default wallet in cookies\r\nfunction changeDefaultWallet(event) {\r\n\tconst idElementClicked = event.target.id;\r\n\tconst newDefaultWalletId = idElementClicked.split('-')[2];\r\n\tconst walletName = gWallets[newDefaultWalletId].nickname;\r\n\r\n\tdocument.cookie = 'defaultWallet=' + walletName;\r\n\t$('.popover__title').html(walletName);\r\n\r\n\tgetWebsiteDeployerSC();\r\n}\r\n\r\nasync function getWallets() {\r\n\taxios\r\n\t\t.get('/mgmt/wallet')\r\n\t\t.then((resp) => {\r\n\t\t\tif (resp) {\r\n\t\t\t\tgWallets = resp.data;\r\n\t\t\t\tfeedWallet(gWallets);\r\n\t\t\t}\r\n\t\t})\r\n\t\t.catch((e) => {\r\n\t\t\terrorAlert(e);\r\n\t\t});\r\n}\r\n\r\nasync function deployWebsiteDeployerSC(password) {\r\n\tlet defaultWallet = getDefaultWallet();\r\n\tconst dnsNameInputValue = document.getElementById('websiteName').value;\r\n\r\n\tif (dnsNameInputValue == '') {\r\n\t\terrorAlert('Input a DNS name');\r\n\t} else {\r\n\t\tdocument.getElementsByClassName('loading')[0].style.display = 'inline-block';\r\n\t\taxios\r\n\t\t\t.put(\r\n\t\t\t\t'/websiteCreator/prepare/',\r\n\t\t\t\t{ url: dnsNameInputValue, nickname: defaultWallet },\r\n\t\t\t\t{\r\n\t\t\t\t\theaders: {\r\n\t\t\t\t\t\tAuthorization: password,\r\n\t\t\t\t\t},\r\n\t\t\t\t}\r\n\t\t\t)\r\n\t\t\t.then((operation) => {\r\n\t\t\t\tdocument.getElementsByClassName('loading')[0].style.display = 'none';\r\n\t\t\t\tsuccessMessage('Contract deployed to address ' + operation.data.address);\r\n\t\t\t\tgetWebsiteDeployerSC();\r\n\t\t\t})\r\n\t\t\t.catch((e) => {\r\n\t\t\t\tdocument.getElementsByClassName('loading')[0].style.display = 'none';\r\n\t\t\t\terrorAlert(e.response.data.code);\r\n\t\t\t});\r\n\t}\r\n}\r\n\r\nfunction tableInsert(resp, count) {\r\n\tconst tBody = document.getElementById('website-deployers-table').getElementsByTagName('tbody')[0];\r\n\tconst row = tBody.insertRow(-1);\r\n\tconst url = 'http://' + resp.name + '.massa/';\r\n\r\n\tconst cell0 = row.insertCell();\r\n\tconst cell1 = row.insertCell();\r\n\tconst cell2 = row.insertCell();\r\n\tconst cell3 = row.insertCell();\r\n\r\n\tcell0.innerHTML = resp.name;\r\n\tcell1.innerHTML = resp.address;\r\n\tcell2.innerHTML = \"<a href='\" + url + \"'>\" + url + '</a>';\r\n\tcell3.innerHTML =\r\n\t\t\"<div><input id='fileid\" +\r\n\t\tcount +\r\n\t\t\"' type='file' hidden/><button id='upload-website\" +\r\n\t\tcount +\r\n\t\t\"' class='primary-button' id='buttonid' type='button' value='Upload MB' >Upload</button><img src='./logo.png' style='display: none' class='massa-logo-spinner loading\" +\r\n\t\tcount +\r\n\t\t\" alt='Massa logo' /></span></div>\";\r\n\r\n\tdocument.getElementById(`upload-website${count}`).addEventListener('click', function () {\r\n\t\tdocument.getElementById(`fileid${count}`).value = null;\r\n\t\tdocument.getElementById(`fileid${count}`).click();\r\n\t});\r\n\r\n\tdocument.getElementById(`fileid${count}`).addEventListener('change', function (evt) {\r\n\t\tlet files = evt.target.files; // get files\r\n\t\tnextFileToUpload = files[0];\r\n\r\n\t\tsetTxType('uploadWebsiteCreator' + count);\r\n\t\t$('#passwordModal').modal('show');\r\n\t});\r\n}\r\n\r\nfunction uploadWebsite(file, count, password) {\r\n\tlet defaultWallet = getDefaultWallet();\r\n\tconst bodyFormData = new FormData();\r\n\tbodyFormData.append('zipfile', file);\r\n\tbodyFormData.append('address', deployers[count].address);\r\n\tbodyFormData.append('nickname', defaultWallet);\r\n\tdocument.getElementsByClassName('loading' + count)[0].style.display = 'inline-block';\r\n\taxios({\r\n\t\turl: `/websiteCreator/upload`,\r\n\t\tmethod: 'POST',\r\n\t\tdata: bodyFormData,\r\n\t\theaders: {\r\n\t\t\t'Content-Type': 'multipart/form-data',\r\n\t\t\tAuthorization: password,\r\n\t\t},\r\n\t})\r\n\t\t.then((operation) => {\r\n\t\t\tdocument.getElementsByClassName('loading' + count)[0].style.display = 'none';\r\n\t\t\tsuccessMessage('Website uploaded to address : ' + operation.data.address);\r\n\t\t})\r\n\t\t.catch((e) => {\r\n\t\t\tdocument.getElementsByClassName('loading' + count)[0].style.display = 'none';\r\n\t\t\terrorAlert(e.response.data.code);\r\n\t\t});\r\n}\r\n"
