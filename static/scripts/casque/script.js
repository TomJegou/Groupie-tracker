import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
console.log("Tout marche")
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera( 100, window.innerWidth / window.innerHeight, 0.1, 1000 );
const renderer = new THREE.WebGLRenderer();
renderer.setSize( window.innerWidth/2, window.innerHeight/2 );
renderer.setPixelRatio(devicePixelRatio);
document.getElementById("object3d").appendChild( renderer.domElement );

const loader = new GLTFLoader();
loader.load(
  "http://127.0.0.1/static/img/casque.gltf",
  function (gltf) {
    scene.add(gltf.scene);
    camera.position.z = 50;
    renderer.render( scene, camera );
  },
  function (xhr) {
    console.log((xhr.loaded / xhr.total) * 100 + "% loaded");
  },
  function (error) {
    console.error(error);
  }
);
// // function animate() {
// //     requestAnimationFrame( animate );
// //     cube.rotation.y += 0.008;
// //     renderer.render( scene, camera );
// // }
// // animate();