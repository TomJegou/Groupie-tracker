import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
import { DRACOLoader } from 'three/examples/jsm/loaders/DRACOLoader.js';

init();

function init() {
    const canvas = document.querySelector('.webgl');
    const scene = new THREE.Scene();
    
    const dracoLoader = new DRACOLoader();
    dracoLoader.setDecoderPath( 'js/libs/draco/gltf/' );
    const loader = new GLTFLoader();
    loader.setDRACOLoader( dracoLoader );

    loader.load('http://localhost/static/img/casque.gltf', function(gltf) {
        const model = gltf.scene;
        model.position.set(1, 1, 0);
        model.scale.set(0.01, 0.01, 0.01);
        scene.add(model);
    }, function(error) {
        console.log("An error occurred: ", error);
    });

    const light = new THREE.DirectionalLight(0xffffff, 1);
    light.position.set(2, 2, 5);
    scene.add(light);

    const sizes = {
        width: window.innerWidth,
        height: window.innerHeight
    };

    const camera = new THREE.PerspectiveCamera(75, sizes.width / sizes.height, 0.1, 100);
    camera.position.set(0, 1, 2);
    scene.add(camera);

    const renderer = new THREE.WebGLRenderer({
        canvas: canvas,
        antialias: true
    });
    renderer.setSize(sizes.width, sizes.height);
    renderer.setPixelRatio(window.devicePixelRatio);
    renderer.render(scene, camera);
}