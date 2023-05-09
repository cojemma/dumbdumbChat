const cubism4Model =
"https://cdn.jsdelivr.net/gh/guansss/pixi-live2d-display/test/assets/haru/haru_greeter_t03.model3.json";

const live2d = PIXI.live2d;
var live2dModel;

const app = new PIXI.Application({
  view: document.getElementById("canvas"),
  autoStart: true,
  resizeTo: window,
  backgroundColor: 0x333333 });

function loadLive2d() {
  (async function main() {
    app.stage.removeChild(live2dModel);
    live2dModel = await live2d.Live2DModel.from(cubism2Model);

    app.stage.addChild(live2dModel);

    const scaleX = innerWidth * 0.4 / live2dModel.width;
    const scaleY = innerHeight * 0.8 / live2dModel.height;

    // fit the window
    live2dModel.scale.set(Math.min(scaleX, scaleY));

    live2dModel.y = innerHeight * 0.1;

    draggable(live2dModel);
    addHitAreaFrames(live2dModel);
    
    addModelControl();

    live2dModel.x = (innerWidth - live2dModel.width) / 2;

    // handle tapping

    live2dModel.on("hit", hitAreas => {
      if (hitAreas.includes("body")) {
        live2dModel.motion("tap_body");
      }

      if (hitAreas.includes("head")) {
        live2dModel.expression();
      }
    });
  })();
};

function draggable(model) {
  model.buttonMode = true;
  model.on("pointerdown", e => {
    model.dragging = true;
    model._pointerX = e.data.global.x - model.x;
    model._pointerY = e.data.global.y - model.y;
  });
  model.on("pointermove", e => {
    if (model.dragging) {
      model.position.x = e.data.global.x - model._pointerX;
      model.position.y = e.data.global.y - model._pointerY;
    }
  });
  model.on("pointerupoutside", () => model.dragging = false);
  model.on("pointerup", () => model.dragging = false);
}

function addFrame(model) {
  const foreground = PIXI.Sprite.from(PIXI.Texture.WHITE);
  foreground.width = model.internalModel.width;
  foreground.height = model.internalModel.height;
  foreground.alpha = 0.2;

  model.addChild(foreground);

  checkbox("Model Frames", checked => foreground.visible = checked);
}

function addHitAreaFrames(model) {
  const hitAreaFrames = new live2d.HitAreaFrames();

  model.addChild(hitAreaFrames);

  checkbox("Hit Area Frames", checked => hitAreaFrames.visible = checked);
}

function addModelControl() {

  checkbox("Montion Control",function (checked) {
    let controlBox = document.getElementById("motionControl");
    controlBox.hidden = !checked;
  });

  checkbox("Expression Control",function (checked) {
    let controlBox = document.getElementById("expressionControl");
    controlBox.hidden = !checked;
  });
}

function checkbox(name, onChange) {
  const id = name.replace(/\W/g, "").toLowerCase();

  let checkbox = document.getElementById(id);

  if (!checkbox) {
    const p = document.createElement("p");
    p.innerHTML = `<input type="checkbox" id="${id}" class="checkbox"> <label for="${id}" class="checkboxLabel">${name}</label>`;

    document.getElementById("control").appendChild(p);
    checkbox = p.firstChild;
  }

  checkbox.addEventListener("change", () => {
    onChange(checkbox.checked);
  });

  onChange(checkbox.checked);
}

function setMotion(move, level) {
  console.log("set motion:", move)
  live2dModel.motion(move, undefined, level);
}

function setExpression(expr, level) {
  console.log("set expression:", expr)
  live2dModel.expression(expr, undefined, level);
}

function changeModel() {
  const select = document.querySelector('#selectModel');
  const selectedOption = select.selectedOptions[0];
  const modelName = selectedOption.innerText;
  let modelResource = document.getElementById("selectModel").value;
  cubism2Model = modelResource;

  fetch('/changeLive2d', {
    method: 'POST', 
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify({ 
      modelName
    })
  })

  loadLive2d();
}

loadLive2d();