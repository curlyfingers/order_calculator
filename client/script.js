async function preload(){
  const response = await fetch('/api/packs', {headers: {'Content-Type': 'application/json'}});

  const {pack_sizes: packSizes = []} = await response.json();
  const container = document.getElementById("available-pack-sizes");
  for(pack of packSizes){
    const div = document.createElement("div");
    div.classList.add("rounded-sm");
    div.classList.add("m-2");
    div.classList.add("px-2");
    div.classList.add("bg-gray-300");
    div.textContent = pack;
    container.appendChild(div);
  }

  const btn = document.getElementById("calculate-btn");
  btn.addEventListener("click", fecthOrderConfiguration);
}

async function fecthOrderConfiguration(event) {
  event.preventDefault();

  const errorContainer = document.getElementById("error");
  const resultContainer = document.getElementById("order-configuration");

  const orderSize = document.getElementById("order-size").value;
  const response = await fetch('/api/order', {
    headers: {'Content-Type': 'application/json'},
    method: 'POST',
    body: JSON.stringify({ order_size: parseInt(orderSize, 10) }),
  });

  const {error, order_configuration: orderConfiguration} = await response.json();
  if(!response.ok){
    resultContainer.parentElement.classList.add("hidden");
    errorContainer.textContent = error;
    errorContainer.classList.remove("hidden");
    return;
  }

  errorContainer.classList.add("hidden");

  resultContainer.innerHTML = "";
  resultContainer.parentElement.classList.remove("hidden");
  for(item of Object.getOwnPropertyNames(orderConfiguration)){
    const div = document.createElement("div");
    div.classList.add("rounded-sm");
    div.classList.add("m-2");
    div.classList.add("px-2");
    div.classList.add("bg-gray-300");
    div.textContent = item;
    const configEntry = document.createElement("div");
    configEntry.classList.add("flex");
    configEntry.classList.add("items-center");
    configEntry.classList.add("flex-row");
    const quantityNode = document.createElement("div");
    quantityNode.textContent = `${orderConfiguration[item]} x`;
    configEntry.appendChild(quantityNode);
    configEntry.appendChild(div);

    resultContainer.appendChild(configEntry)
  }
}

preload();
