import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera( 100, window.innerWidth / window.innerHeight, 0.1, 1000 );
const renderer = new THREE.WebGLRenderer();
renderer.setSize( window.innerWidth/2, window.innerHeight/2 );
renderer.setPixelRatio(devicePixelRatio);
renderer.setClearColor(new THREE.Color("rgb(3, 165, 209)"));
document.getElementById("object3d").appendChild( renderer.domElement );

const light = new THREE.PointLight( 0xffffff, 1, 100 );
light.position.set( 0, 0, 10 );
scene.add( light );

function getGradientCanvas() {
    const canvas = document.createElement('canvas');
    canvas.width = 64;
    canvas.height = 64;
    const context = canvas.getContext('2d');
    const gradient = context.createLinearGradient(0, 0, 0, 64);
    gradient.addColorStop(0, 'rgb(3, 165, 209)');
    gradient.addColorStop(1, 'white');
    context.fillStyle = gradient;
    context.fillRect(0, 0, 64, 64);
    return canvas;
}

const gradientTexture = new THREE.CanvasTexture(getGradientCanvas());
gradientTexture.wrapS = THREE.RepeatWrapping;
gradientTexture.wrapT = THREE.RepeatWrapping;
scene.background = gradientTexture;

let object;

const loader = new GLTFLoader();
loader.load(
    "http://127.0.0.1/static/img/casque.gltf",
    function (gltf) {
        object = gltf.scene;
        object.position.x = -2; // déplacer l'objet sur la gauche
        object.rotateX(1.5);
        scene.add(object);
        camera.position.z = 10;
        camera.position.x = 5;
        camera.position.y = 10;
        camera.lookAt(object.position);
        renderer.render( scene, camera );
    },
    function (xhr) {
        console.log((xhr.loaded / xhr.total) * 100 + "% loaded");
    },
    function (error) {
        console.error(error);
    }
    );

const controls = new OrbitControls(camera, renderer.domElement);
controls.autoRotate = true;
controls.update();

function animate() {
    requestAnimationFrame( animate );
    controls.update();
    renderer.render( scene, camera );
}
animate();