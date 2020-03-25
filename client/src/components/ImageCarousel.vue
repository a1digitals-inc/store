<template>
    <div>
        <img :src="imageSource" />
    </div>
</template>

<script>
export default {
    name: "ImageCarousel",
    props: ["images"],
    data () {
        return {
            source: 0
        }
    },
    created () {
        setInterval(() => {
            this.source++
            if (this.source >= this.images.length) {
                this.source = 0
            }
        }, 5000)
    },
    computed: {
        imageSource() { 
            return this.images[this.source]
        }
    }
}
</script>
