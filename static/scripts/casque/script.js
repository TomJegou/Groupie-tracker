import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';

const scene = new THREE.Scene();
scene.background = null
const ambientLight = new THREE.AmbientLight( 0xffffff, 0.5 );
scene.add( ambientLight );

const camera = new THREE.PerspectiveCamera( 100, window.innerWidth / window.innerHeight, 0.1, 1000 );
const renderer = new THREE.WebGLRenderer({ alpha: true });
renderer.setSize( window.innerWidth/2, window.innerHeight/2 );
renderer.setPixelRatio(devicePixelRatio);
document.getElementById("object3d").appendChild( renderer.domElement );

const light = new THREE.PointLight( 0xffffff, 1, 100 );
light.position.set( 0, 0, 10 );
scene.add( light );

let object;

const loader = new GLTFLoader();
loader.load(
    "static/img/casque.gltf",
    function (gltf) {
        object = gltf.scene;
        object.traverse(function(child) {
            if (child instanceof THREE.Mesh) {
                child.material.transparent = true;
            }
        });
        object.position.x = -2;
        object.rotateX(Math.PI - 1);
        scene.add(object);
        camera.position.z = 10;
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