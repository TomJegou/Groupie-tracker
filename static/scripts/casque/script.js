import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
console.log("Tout marche")
const scene = new THREE.Scene();
scene.background = new THREE.Color("#03a5d1");
const camera = new THREE.PerspectiveCamera( 100, window.innerWidth / window.innerHeight, 0.1, 1000 );
const renderer = new THREE.WebGLRenderer();
renderer.setSize( window.innerWidth/2, window.innerHeight/2 );
renderer.setPixelRatio(devicePixelRatio);
document.getElementById("object3d").appendChild( renderer.domElement );

const light = new THREE.PointLight( 0xffffff, 1, 100 );
light.position.set( 0, 0, 10 );
scene.add( light );

const loader = new GLTFLoader();
loader.load(
  "http://127.0.0.1/static/img/casque.gltf",
  function (gltf) {
    const object = gltf.scene;
    object.rotateX(1);
    scene.add(object);
    camera.position.z = 10;
    camera.position.x = 5;
    renderer.render( scene, camera );
  },
  function (xhr) {
    console.log((xhr.loaded / xhr.total) * 100 + "% loaded");
  },
  function (error) {
    console.error(error);
  }
);