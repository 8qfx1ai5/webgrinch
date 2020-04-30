package main

import (
	"fmt"
	"log"

	"github.com/8qfx1ai5/viewcrypt/internal/encodehtml"
)

func main() {
	in := "<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Molestiae minus at aut illo esse unde nemo sint reprehenderit et veritatis qui vel aspernatur sunt, explicabo consequuntur obcaecati similique suscipit dicta! <a rel=\"noreferrer noopener\" href=\"https://example.com/foo/bar\" target=\"_blank\">Link</a> Lorem, ipsum dolor sit amet consectetur adipisicing elit. Laborum quisquam nobis doloremque, est ut veritatis.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Debitis autem aliquam temporibus saepe, nisi suscipit eius accusamus fuga nesciunt porro sequi qui doloremque voluptates doloribus facere maxime vero deleniti similique aut sed nostrum placeat sapiente cumque molestias. Accusantium, possimus sit! Fuga laborum non, quod repellendus inventore iusto commodi nihil, magnam culpa saepe, expedita quibusdam ratione deserunt. Vel voluptas possimus expedita et molestias, dolor ratione! Voluptas facere aperiam labore fugit? Modi nulla ducimus esse rem alias voluptatem eum praesentium consectetur placeat atque omnis architecto, maiores aperiam nihil fugiat magni debitis sint beatae blanditiis quidem harum molestias recusandae! Nostrum asperiores porro dicta nisi debitis quas commodi eaque expedita nam eum animi, quisquam vero dolore officia reiciendis ab magni impedit praesentium voluptatibus deleniti! Vitae quidem consequatur dicta ipsam in ipsum reprehenderit quae accusamus itaque. Architecto aliquam mollitia vel vero, veritatis magnam tempora illo, sint, earum minus explicabo consequuntur similique. Nemo blanditiis expedita nulla?</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Unde, numquam! Voluptatibus, earum, veritatis molestiae assumenda totam nemo accusantium facere labore repellat laudantium deleniti ut distinctio recusandae necessitatibus consequuntur quibusdam! Atque eveniet hic voluptas eos blanditiis dicta explicabo eligendi quis rem a? Voluptate quod in dolorem sequi beatae consequatur laudantium magni.</p>\n<!-- /wp:paragraph -->"
	keyFrom := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	keyTo := "aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO"
	cssClass := "vc"
	out, err := encodehtml.Run(in, keyFrom, keyTo, cssClass)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
