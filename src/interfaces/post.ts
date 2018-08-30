import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IPost extends Typegoose {

	@prop({ required: true })
        schema_version: number;

	@prop({ required: true })
	id: string;

	@prop({ required: true })
	account: string;

	@prop({ required: true })
	content: string;

	@prop({ required: true })
	created_at: number;

	@prop()
	multimedia?: IAttachment;

	@prop()
	tags?: Array<string>;

	@prop({ required: true })
	likes_count: number;

	@prop({ required: true })
	comments_count: number;

	@prop({ required: true })
	shares_count: number;

	@prop()
	likes?: Array<string>;

	@prop()
	comments?: Array<IComment>;

	@prop()
	shares?: Array<string>;

	@prop({ required: true })
	visibility: string;

}
