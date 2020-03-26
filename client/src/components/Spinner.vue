<template>
    <div>
    </div>
</template>

<style scoped>
div {
    margin: auto;
    pointer-events: none;
    width: 1em;
    height: 1em;
    border: 0.4em solid transparent;
    border-top-color: black;
    border-radius: 50%;
    animation: loadingspin 1s linear infinite;
}

@keyframes loadingspin{
    100% {
        transform: rotate(360deg)
    }
}
</style>
