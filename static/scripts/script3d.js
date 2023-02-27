import * as THREE from 'three';


let container;
let camera;
let scene;
let casque;

function init(){
    container = document.querySelector('.scene')
    scene = new THREE.Scene('.scene');
    const fov = 35;
    const aspect = container.clientWidth / container.clientHeight;
    const near = 0.1;
    const far = 500;
    camera = new THREE.PerspectiveCamera(fov,aspect,near,far)

    camera.position.set(-50, 40, 350);

    renderer = new THREE.WebGLRenderer({antialias:true, alpha: true});
    renderer.setSize(container.clientWidth,container.clientHeight);
    renderer.setPixelRatio(window.devicePixelRatio)
    container.appendChild(renderer.domElement);

    let loader = new THREE.GLTFLoader();
    loader.load("http://localhost:8080/static/img/casque.gltf", function(gltf){
        scene.add(gltf.scene);
    });

}
init()