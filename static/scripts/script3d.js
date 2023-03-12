import * as THREE from './node_modules/three/build/three.module.js';
import * as gltf from './node_modules/three/examples/jsm/loaders/GLTFLoader.js';
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 0.1, 1000 );
const renderer = new THREE.WebGLRenderer();
renderer.setSize( window.innerWidth/2, window.innerHeight/2 );
renderer.setPixelRatio(devicePixelRatio);
document.getElementById("object3d").appendChild( renderer.domElement );

const loader = new gltf.GLTFLoader();
loader.load( "../img/casque.gltf", function (gltf) {
    scene.add(gltf.scene)
} );
camera.position.z = 10;
// function animate() {
//     requestAnimationFrame( animate );
//     cube.rotation.y += 0.008;
//     renderer.render( scene, camera );
// }
animate();