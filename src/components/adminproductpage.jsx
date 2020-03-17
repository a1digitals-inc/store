import React from "react";

class AdminProductPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            message: "",
            id: "",
            thumbnail: "",
            images: [],
            data: {
                name: "",
                identifier: "",
                public: false,
                thumbnail: "",
                images: [],
                description: "",
                price: "",
                discount: 1
            }
        };
        this.submit = this.submit.bind(this)
        this.thumbnailUpload = this.thumbnailUpload.bind(this)
        this.imageUpload = this.imageUpload.bind(this)
    }

    componentDidMount() {
        if (window.location.pathname != "/dashboard/product") {
            var id = window.location.pathname.split("/").pop();
            fetch("/api/admin/product/" + id).then(res => res.json())
                .then(
                    (result) => {
                        document.title = result.name;
                        this.setState({
                            isLoaded: true,
                            data: result,
                            thumbnail: result.thumbnail,
                            images: result.images,
                            id
                        });
                    },
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        });
                    }
                );
        } else {
            this.setState({
                isLoaded: true,
            });
        }
        return
    }

    submit(event) {
        event.preventDefault()
        const data = new FormData(event.target)
        const options = {
            method: "PUT",
            body: data
        };
        fetch("/api/product", options)
            .then(res => res.json())
            .then(
                (result) => {
                    if (result.message) {
                        this.setState({
                            message: result.message,
                        });
                    }
                },
                (error) => {
                    this.setState({
                        message: error 
                    });
                }
            )
    }

    thumbnailUpload(event) {
        this.setState({
            thumbnail: URL.createObjectURL(event.target.files[0])
        });
    }

    imageUpload(event) {
        var images = [];
        for (var i = 0, l; i < event.target.files.length; i++) {
            images.push(URL.createObjectURL(event.target.files[i]));
        }
        this.setState({
            images
        });
    }

    render() {
        const {error, isLoaded, id, thumbnail, images, data} = this.state;
        if (error) {
            return <div className="text-center">Error {error.message}</div>;
        } else if (!isLoaded) {
            return (
                <div className="container-fluid text-center">
                    <h1 className="m-5"><a href="/dashboard/products">Dashboard</a></h1>
                    <div className="spinner-border"></div>
                </div>
            );
        } else {
            return (
                <div className="container-fluid">
                    <h1 className="m-5 text-center"><a href="/dashboard/products">Dashboard</a></h1>
                    <form className="col-sm-10 m-auto p-2" onSubmit={this.submit}>
                        {this.state.message && <div className="fade-in alert alert-secondary">{this.state.message}</div>}
                        <div className="form-group">
                            <label>Name</label>
                            <input className="form-control" type="text" name="name" defaultValue={data.name}/>
                        </div>
                        <div className="form-group">
                            <label>Identifier</label>
                            <input className="form-control" type="text" name="identifier" defaultValue={id} />
                            <small className="form-text text-muted">Unique identifier: No spaces, No special characters</small>
                        </div>
                        <div className="form-check">
                            <input className="form-check-input" type="checkbox" name="public" defaultChecked={data.public}/>
                            <label>Public</label>
                        </div>
                        <div className="form-group">
                            <label>Thumbnail</label>
                            <input className="form-control-file" type="file" name="thumbnail" onChange={this.thumbnailUpload} />
                            <img src={thumbnail} className="mini-image my-2" />
                        </div>
                        <div className="form-group">
                            <label>Images</label>
                            <input className="form-control-file" type="file" name="images" multiple onChange={this.imageUpload} />
                            {images && images.map((image, i) => (
                                <img key={i} src={image} className="mini-image my-2 mr-2"/>
                            ))}
                        </div>
                        <div className="form-group">
                            <label>Description</label>
                            <textarea className="form-control" rows="3" name="description" defaultValue={data.description} />
                        </div>
                        <div className="form-group">
                            <label>Price</label>
                            <input className="form-control" type="number" name="price" defaultValue={data.price} />
                        </div>
                        <div className="form-group">
                            <label>Discount</label>
                            <input className="form-control" type="number" step="0.01" min="0" max="1" name="discount" defaultValue={data.discount} />
                            <small className="form-text text-muted">Discount multiplier applied to price</small>
                        </div>
                        <button type="submit" className="btn btn-primary float-right">Submit</button>
                    </form>
                </div>
            );    
        }
    }
}

export default AdminProductPage;
