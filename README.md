![swift logo](https://eclipsesource.com/wp-content/uploads/2014/06/Apple_Swift_Logo-150x150.png)
# Genwift
Simple utility written on Go to create templates for modules in swift 
# Usage 
run ```go build``` and copy binary file to your root directory(for example into your user's directory).

For create MVP module with assembly, presenter, view controller, protocols for view and presenter you need to write module name as argument.
```~/genwift Login```

# Files structure 
```~/genwift Login```

#### LoginViewController.swift 
```swift 
class LoginViewController: LoginInput { 
	<#T##ViewController's code#> 
}
```

#### LoginPresenter.swift 

```swift 
class LoginPresenter: LoginOutput {
	weak var view:LoginInput? 
	<#T##Presenter's code#> 
}
init(view: Login) {
	self.view = view 
}
```

#### LoginInput.swift 
```swift 
protocol LoginInput { 
	<#T##Input's protocol code#> 
}
```

#### LoginOutput.swift 
```swift
protocol LoginOutput { 
	<#T##Output's protocol code#> 
}
```

#### LoginAssembly.swift 

```swift 
class LoginAssembly { 
 class func configureModule() -> LoginViewController {
	let view = LoginViewController()
	let presenter = LoginPresenter(view: view)
	view.presenter = presenter
	return view
	}
}
```
