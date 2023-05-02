const cubism2Model =
"https://cdn.jsdelivr.net/gh/guansss/pixi-live2d-display/test/assets/shizuku/shizuku.model.json";
const cubism4Model =
"https://cdn.jsdelivr.net/gh/guansss/pixi-live2d-display/test/assets/haru/haru_greeter_t03.model3.json";

const live2d = PIXI.live2d;
var shizuku;
var haru;

(async function main() {
  const app = new PIXI.Application({
    view: document.getElementById("canvas"),
    autoStart: true,
    resizeTo: window,
    backgroundColor: 0x333333 });

  const models = await Promise.all([
    live2d.Live2DModel.from(cubism2Model)]);

  models.forEach(model => {
    app.stage.addChild(model);

    const scaleX = innerWidth * 0.4 / model.width;
    const scaleY = innerHeight * 0.8 / model.height;

    // fit the window
    model.scale.set(Math.min(scaleX, scaleY));

    model.y = innerHeight * 0.1;

    draggable(model);
    addHitAreaFrames(model);
  });
  
  addModelControl();

  const model2 = models[0];
  shizuku = models[0];

  model2.x = (innerWidth - model2.width) / 2;

  // handle tapping

  model2.on("hit", hitAreas => {
    if (hitAreas.includes("body")) {
      model2.motion("tap_body");
    }

    if (hitAreas.includes("head")) {
      model2.expression();
    }
  });
})();

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
    p.innerHTML = `<input type="checkbox" id="${id}"> <label for="${id}">${name}</label>`;

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
  shizuku.motion(move, undefined, level);
}

function setExpression(expr, level) {
  console.log("set expression:", expr)
  shizuku.expression(expr, undefined, level);
}
