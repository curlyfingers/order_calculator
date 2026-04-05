async function preload(){
  const response = await fetch('/api/packs', {headers: {'Content-Type': 'application/json'}});

  const {pack_sizes: packSizes = []} = await response.json();
  const container = document.getElementById("available-pack-sizes");
  const colors = ['red-500', 'blue-500', 'green-500', 'yellow-500', 'purple-500','orange-500','yellow-500'];
  for(pack of packSizes){
    const randomColor = colors[Math.floor(Math.random() * colors.length)];
    const div = document.createElement("div");
    div.classList.add("border-4");
    div.classList.add("rounded-xl");
    div.classList.add("m-2");
    div.classList.add("px-2");
    div.classList.add("border-"+randomColor);
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

  const body = await response.json();
  console.log(body)
  if(!response.ok){
    const error = body.error;
    errorContainer.textContent = error;
    errorContainer.classList.remove("hidden");
    return;
  }

  errorContainer.classList.add("hidden");

  const orderConfiguration = body.order_configuration;
  resultContainer.innerHTML = "";
  resultContainer.parentElement.classList.remove("hidden");
  for(item of Object.getOwnPropertyNames(orderConfiguration)){
    const child = document.createElement("div");
    child.textContent = `${orderConfiguration[item]} x ${item}`
    resultContainer.appendChild(child)
  }
}

preload();
