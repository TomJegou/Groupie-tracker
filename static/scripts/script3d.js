const sceneContainer = document.getElementById('scene-container');
const renderer = new THREE.WebGLRenderer();
sceneContainer.appendChild(renderer.domElement);

const scene = new THREE.Scene();

const camera = new THREE.PerspectiveCamera()
const ambientLight = new THREE.AmbientLight();